package cmdutils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/printers"
	"github.com/weaveworks/eksctl/pkg/utils/kubeconfig"
	"github.com/weaveworks/eksctl/pkg/version"
)

// IncompatibleFlags is a common substring of an error message
const IncompatibleFlags = "cannot be used at the same time"

// NewVerbCmd defines a standard verb command
func NewVerbCmd(use, short, long string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(c *cobra.Command, _ []string) {
			if err := c.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},
	}
}

// AddPreRun chains cmd.PreRun handlers, as cobra only allows one, so we don't
// accidentally override one we registered earlier
func AddPreRun(cmd *cobra.Command, newFn func(cmd *cobra.Command, args []string)) {
	currentFn := cmd.PreRun
	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		if currentFn != nil {
			currentFn(cmd, args)
		}
		newFn(cmd, args)
	}
}

// LogIntendedAction calls logger.Info with appropriate prefix
func LogIntendedAction(plan bool, msgFmt string, args ...interface{}) {
	prefix := "will "
	if plan {
		prefix = "(plan) would "
	}
	logger.Info(prefix+msgFmt, args...)
}

// LogCompletedAction calls logger.Success with appropriate prefix
func LogCompletedAction(plan bool, msgFmt string, args ...interface{}) {
	prefix := ""
	if plan {
		prefix = "(plan) would have "
	}
	logger.Success(prefix+msgFmt, args...)
}

// LogPlanModeWarning will log a message to inform user that they are in plan-mode
func LogPlanModeWarning(plan bool) {
	if plan {
		logger.Warning("no changes were applied, run again with '--approve' to apply the changes")
	}
}

// LogRegionAndVersionInfo will log the selected region and build version
func LogRegionAndVersionInfo(meta *api.ClusterMeta) {
	if meta != nil {
		logger.Info("eksctl version %s", version.GetVersion())
		logger.Info("using region %s", meta.Region)
	}
}

// AddApproveFlag adds common `--approve` flag
func AddApproveFlag(fs *pflag.FlagSet, cmd *Cmd) {
	approve := fs.Bool("approve", !cmd.Plan, "Apply the changes")
	AddPreRun(cmd.CobraCommand, func(cobraCmd *cobra.Command, args []string) {
		if cobraCmd.Flag("approve").Changed {
			cmd.Plan = !*approve
		}
	})
}

// GetNameArg tests to ensure there is only 1 name argument
func GetNameArg(args []string) string {
	if len(args) > 1 {
		logger.Critical("only one argument is allowed to be used as a name")
		os.Exit(1)
	}
	if len(args) == 1 {
		return strings.TrimSpace(args[0])
	}
	return ""
}

// AddCommonFlagsForAWS adds common flags for api.ProviderConfig
func AddCommonFlagsForAWS(cmd *Cmd, p *api.ProviderConfig, addCfnOptions bool) {
	cmd.FlagSetGroup.InFlagSet("AWS client", func(fs *pflag.FlagSet) {
		fs.StringVarP(&p.Profile.Name, "profile", "p", "", "AWS credentials profile to use (defaults to the value of the AWS_PROFILE environment variable)")
		if addCfnOptions {
			fs.StringVar(&p.CloudFormationRoleARN, "cfn-role-arn", "", "IAM role used by CloudFormation to call AWS API on your behalf")
			fs.BoolVar(&p.CloudFormationDisableRollback, "cfn-disable-rollback", false, "for debugging: If a stack fails, do not roll it back. Be careful, this may lead to unintentional resource consumption!")
		}
	})

	AddPreRun(cmd.CobraCommand, func(c *cobra.Command, args []string) {
		if !c.Flag("profile").Changed {
			if val, ok := os.LookupEnv("AWS_PROFILE"); ok {
				p.Profile = api.Profile{
					Name:           val,
					SourceIsEnvVar: true,
				}
			}
		}
	})
}

// AddTimeoutFlagWithValue configures the timeout flag with the provided value.
func AddTimeoutFlagWithValue(fs *pflag.FlagSet, p *time.Duration, value time.Duration) {
	fs.DurationVar(p, "timeout", value, "maximum waiting time for any long-running operation")
}

// AddTimeoutFlag configures the timeout flag.
func AddTimeoutFlag(fs *pflag.FlagSet, p *time.Duration) {
	AddTimeoutFlagWithValue(fs, p, api.DefaultWaitTimeout)
}

// AddClusterFlag adds a common --cluster flag for cluster name.
// Use this for commands whose principal resource is *not* a cluster.
func AddClusterFlag(fs *pflag.FlagSet, meta *api.ClusterMeta) {
	fs.StringVarP(&meta.Name, "cluster", "c", "", "EKS cluster name")
}

// AddClusterFlagWithDeprecated adds a common --cluster flag for
// cluster name as well as a deprecated --name flag.
// Use AddClusterFlag() for new commands.
func AddClusterFlagWithDeprecated(fs *pflag.FlagSet, meta *api.ClusterMeta) {
	AddClusterFlag(fs, meta)
	fs.StringVarP(&meta.Name, "name", "n", "", "EKS cluster name")
	_ = fs.MarkDeprecated("name", "use --cluster")
}

// ClusterNameFlag returns the flag to use for the cluster name
// taking the principal resource into account.
func ClusterNameFlag(cmd *Cmd) string {
	if cmd.CobraCommand.Use == "cluster" {
		return "--name"
	}
	return "--cluster"
}

// AddRegionFlag adds common --region flag
func AddRegionFlag(fs *pflag.FlagSet, p *api.ProviderConfig) {
	fs.StringVarP(&p.Region, "region", "r", "", "AWS region. Defaults to the value set in your AWS config (~/.aws/config)")
}

// AddVersionFlag adds common --version flag
func AddVersionFlag(fs *pflag.FlagSet, meta *api.ClusterMeta, extraUsageInfo string) {
	usage := fmt.Sprintf("Kubernetes version (valid options: %s)", strings.Join(api.SupportedVersions(), ", "))
	if extraUsageInfo != "" {
		usage = fmt.Sprintf("%s [%s]", usage, extraUsageInfo)
	}
	fs.StringVar(&meta.Version, "version", meta.Version, usage)
}

// AddWaitFlag adds common --wait flag
func AddWaitFlag(fs *pflag.FlagSet, wait *bool, description string) {
	AddWaitFlagWithFullDescription(fs, wait, fmt.Sprintf("wait for %s before exiting", description))
}

// AddWaitFlagWithFullDescription adds common --wait flag
func AddWaitFlagWithFullDescription(fs *pflag.FlagSet, wait *bool, description string) {
	fs.BoolVarP(wait, "wait", "w", *wait, description)
}

// AddUpdateAuthConfigMap adds common --update-auth-configmap flag
func AddUpdateAuthConfigMap(fs *pflag.FlagSet, description string) *bool {
	return fs.Bool("update-auth-configmap", true, description)
}

// AddSubnetIDs adds common --subnet-ids flag
func AddSubnetIDs(fs *pflag.FlagSet, subnetIDs *[]string, description string) {
	fs.StringSliceVar(subnetIDs, "subnet-ids", nil, description)
}

// AddCommonFlagsForKubeconfig adds common flags for controlling how output kubeconfig is written
func AddCommonFlagsForKubeconfig(fs *pflag.FlagSet, outputPath, authenticatorRoleARN *string, setContext, autoPath *bool, exampleName string) {
	fs.StringVar(outputPath, "kubeconfig", kubeconfig.DefaultPath(), "path to write kubeconfig (incompatible with --auto-kubeconfig)")
	fs.StringVar(authenticatorRoleARN, "authenticator-role-arn", "", "AWS IAM role to assume for authenticator")
	fs.BoolVar(setContext, "set-kubeconfig-context", true, "if true then current-context will be set in kubeconfig; if a context is already set then it will be overwritten")
	fs.BoolVar(autoPath, "auto-kubeconfig", false, fmt.Sprintf("save kubeconfig file by cluster name, e.g. %q", kubeconfig.AutoPath(exampleName)))
}

// AddCommonFlagsForGetCmd adds common flags for get commands.
func AddCommonFlagsForGetCmd(fs *pflag.FlagSet, chunkSize *int, outputMode *printers.Type) {
	fs.IntVar(chunkSize, "chunk-size", 100, "return large lists in chunks rather than all at once, pass 0 to disable")
	fs.StringVarP(outputMode, "output", "o", "table", "specifies the output format (valid option: table, json, yaml)")
}

// AddStringToStringVarPFlag is a wrapper that prefixes the description of the flag for consistency
func AddStringToStringVarPFlag(fs *pflag.FlagSet, p *map[string]string, name, shorthand string, value map[string]string, usage string) {
	fs.StringToStringVarP(p, name, shorthand, value, fmt.Sprintf(`%s. List of comma separated KV pairs "k1=v1,k2=v2"`, usage))
}

// ErrUnsupportedRegion is a common error message
func ErrUnsupportedRegion(provider *api.ProviderConfig) error {
	return fmt.Errorf("--region=%s is not supported - use one of: %s", provider.Region, strings.Join(api.SupportedRegions(), ", "))
}

// ErrClusterFlagAndArg wraps ErrFlagAndArg() by passing in the
// proper flag name.
func ErrClusterFlagAndArg(cmd *Cmd, nameFlag, nameArg string) error {
	return ErrFlagAndArg(ClusterNameFlag(cmd), nameFlag, nameArg)
}

// ErrFlagAndArg may be used to err for options that can be given
// as flags /and/ arg but only one is allowed to be used.
func ErrFlagAndArg(kind, flag, arg string) error {
	return fmt.Errorf("%s=%s and argument %s %s", kind, flag, arg, IncompatibleFlags)
}

// ErrMustBeSet is a common error message
func ErrMustBeSet(pathOrFlag string) error {
	return fmt.Errorf("%s must be set", pathOrFlag)
}

// ErrCannotUseWithConfigFile is a common error message
func ErrCannotUseWithConfigFile(what string) error {
	return fmt.Errorf("cannot use %s when --config-file/-f is set", what)
}

// ErrUnsupportedManagedFlag reports unsupported flags for Managed Nodegroups
func ErrUnsupportedManagedFlag(flag string) error {
	return fmt.Errorf("%s is not supported for Managed Nodegroups (--managed=true)", flag)
}

// ErrUnsupportedNameArg reports unsupported usage of `name` argument
func ErrUnsupportedNameArg() error {
	return errors.New("name argument is not supported")
}
