package resource

import (
	context "context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	github_com_kyverno_kyverno_api_kyverno_v1alpha2 "github.com/kyverno/kyverno/api/kyverno/v1alpha2"
	github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v1alpha2"
	"github.com/kyverno/kyverno/pkg/metrics"
	"github.com/kyverno/kyverno/pkg/tracing"
	"go.uber.org/multierr"
	k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s_io_apimachinery_pkg_types "k8s.io/apimachinery/pkg/types"
	k8s_io_apimachinery_pkg_watch "k8s.io/apimachinery/pkg/watch"
)

func WithLogging(inner github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface, logger logr.Logger) github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface {
	return &withLogging{inner, logger}
}

func WithMetrics(inner github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface, recorder metrics.Recorder) github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface {
	return &withMetrics{inner, recorder}
}

func WithTracing(inner github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface, client, kind string) github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface {
	return &withTracing{inner, client, kind}
}

type withLogging struct {
	inner  github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface
	logger logr.Logger
}

func (c *withLogging) Create(arg0 context.Context, arg1 *github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.CreateOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	start := time.Now()
	logger := c.logger.WithValues("operation", "Create")
	ret0, ret1 := c.inner.Create(arg0, arg1, arg2)
	if err := multierr.Combine(ret1); err != nil {
		logger.Error(err, "Create failed", "duration", time.Since(start))
	} else {
		logger.Info("Create done", "duration", time.Since(start))
	}
	return ret0, ret1
}
func (c *withLogging) Delete(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions) error {
	start := time.Now()
	logger := c.logger.WithValues("operation", "Delete")
	ret0 := c.inner.Delete(arg0, arg1, arg2)
	if err := multierr.Combine(ret0); err != nil {
		logger.Error(err, "Delete failed", "duration", time.Since(start))
	} else {
		logger.Info("Delete done", "duration", time.Since(start))
	}
	return ret0
}
func (c *withLogging) DeleteCollection(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) error {
	start := time.Now()
	logger := c.logger.WithValues("operation", "DeleteCollection")
	ret0 := c.inner.DeleteCollection(arg0, arg1, arg2)
	if err := multierr.Combine(ret0); err != nil {
		logger.Error(err, "DeleteCollection failed", "duration", time.Since(start))
	} else {
		logger.Info("DeleteCollection done", "duration", time.Since(start))
	}
	return ret0
}
func (c *withLogging) Get(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.GetOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	start := time.Now()
	logger := c.logger.WithValues("operation", "Get")
	ret0, ret1 := c.inner.Get(arg0, arg1, arg2)
	if err := multierr.Combine(ret1); err != nil {
		logger.Error(err, "Get failed", "duration", time.Since(start))
	} else {
		logger.Info("Get done", "duration", time.Since(start))
	}
	return ret0, ret1
}
func (c *withLogging) List(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReportList, error) {
	start := time.Now()
	logger := c.logger.WithValues("operation", "List")
	ret0, ret1 := c.inner.List(arg0, arg1)
	if err := multierr.Combine(ret1); err != nil {
		logger.Error(err, "List failed", "duration", time.Since(start))
	} else {
		logger.Info("List done", "duration", time.Since(start))
	}
	return ret0, ret1
}
func (c *withLogging) Patch(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_types.PatchType, arg3 []uint8, arg4 k8s_io_apimachinery_pkg_apis_meta_v1.PatchOptions, arg5 ...string) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	start := time.Now()
	logger := c.logger.WithValues("operation", "Patch")
	ret0, ret1 := c.inner.Patch(arg0, arg1, arg2, arg3, arg4, arg5...)
	if err := multierr.Combine(ret1); err != nil {
		logger.Error(err, "Patch failed", "duration", time.Since(start))
	} else {
		logger.Info("Patch done", "duration", time.Since(start))
	}
	return ret0, ret1
}
func (c *withLogging) Update(arg0 context.Context, arg1 *github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.UpdateOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	start := time.Now()
	logger := c.logger.WithValues("operation", "Update")
	ret0, ret1 := c.inner.Update(arg0, arg1, arg2)
	if err := multierr.Combine(ret1); err != nil {
		logger.Error(err, "Update failed", "duration", time.Since(start))
	} else {
		logger.Info("Update done", "duration", time.Since(start))
	}
	return ret0, ret1
}
func (c *withLogging) Watch(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) (k8s_io_apimachinery_pkg_watch.Interface, error) {
	start := time.Now()
	logger := c.logger.WithValues("operation", "Watch")
	ret0, ret1 := c.inner.Watch(arg0, arg1)
	if err := multierr.Combine(ret1); err != nil {
		logger.Error(err, "Watch failed", "duration", time.Since(start))
	} else {
		logger.Info("Watch done", "duration", time.Since(start))
	}
	return ret0, ret1
}

type withMetrics struct {
	inner    github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface
	recorder metrics.Recorder
}

func (c *withMetrics) Create(arg0 context.Context, arg1 *github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.CreateOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	defer c.recorder.Record("create")
	return c.inner.Create(arg0, arg1, arg2)
}
func (c *withMetrics) Delete(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions) error {
	defer c.recorder.Record("delete")
	return c.inner.Delete(arg0, arg1, arg2)
}
func (c *withMetrics) DeleteCollection(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) error {
	defer c.recorder.Record("delete_collection")
	return c.inner.DeleteCollection(arg0, arg1, arg2)
}
func (c *withMetrics) Get(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.GetOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	defer c.recorder.Record("get")
	return c.inner.Get(arg0, arg1, arg2)
}
func (c *withMetrics) List(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReportList, error) {
	defer c.recorder.Record("list")
	return c.inner.List(arg0, arg1)
}
func (c *withMetrics) Patch(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_types.PatchType, arg3 []uint8, arg4 k8s_io_apimachinery_pkg_apis_meta_v1.PatchOptions, arg5 ...string) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	defer c.recorder.Record("patch")
	return c.inner.Patch(arg0, arg1, arg2, arg3, arg4, arg5...)
}
func (c *withMetrics) Update(arg0 context.Context, arg1 *github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.UpdateOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	defer c.recorder.Record("update")
	return c.inner.Update(arg0, arg1, arg2)
}
func (c *withMetrics) Watch(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) (k8s_io_apimachinery_pkg_watch.Interface, error) {
	defer c.recorder.Record("watch")
	return c.inner.Watch(arg0, arg1)
}

type withTracing struct {
	inner  github_com_kyverno_kyverno_pkg_client_clientset_versioned_typed_kyverno_v1alpha2.AdmissionReportInterface
	client string
	kind   string
}

func (c *withTracing) Create(arg0 context.Context, arg1 *github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.CreateOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "Create"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("Create"),
	)
	defer span.End()
	arg0 = ctx
	ret0, ret1 := c.inner.Create(arg0, arg1, arg2)
	tracing.SetSpanStatus(span, ret1)
	return ret0, ret1
}
func (c *withTracing) Delete(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions) error {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "Delete"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("Delete"),
	)
	defer span.End()
	arg0 = ctx
	ret0 := c.inner.Delete(arg0, arg1, arg2)
	tracing.SetSpanStatus(span, ret0)
	return ret0
}
func (c *withTracing) DeleteCollection(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) error {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "DeleteCollection"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("DeleteCollection"),
	)
	defer span.End()
	arg0 = ctx
	ret0 := c.inner.DeleteCollection(arg0, arg1, arg2)
	tracing.SetSpanStatus(span, ret0)
	return ret0
}
func (c *withTracing) Get(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.GetOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "Get"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("Get"),
	)
	defer span.End()
	arg0 = ctx
	ret0, ret1 := c.inner.Get(arg0, arg1, arg2)
	tracing.SetSpanStatus(span, ret1)
	return ret0, ret1
}
func (c *withTracing) List(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReportList, error) {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "List"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("List"),
	)
	defer span.End()
	arg0 = ctx
	ret0, ret1 := c.inner.List(arg0, arg1)
	tracing.SetSpanStatus(span, ret1)
	return ret0, ret1
}
func (c *withTracing) Patch(arg0 context.Context, arg1 string, arg2 k8s_io_apimachinery_pkg_types.PatchType, arg3 []uint8, arg4 k8s_io_apimachinery_pkg_apis_meta_v1.PatchOptions, arg5 ...string) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "Patch"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("Patch"),
	)
	defer span.End()
	arg0 = ctx
	ret0, ret1 := c.inner.Patch(arg0, arg1, arg2, arg3, arg4, arg5...)
	tracing.SetSpanStatus(span, ret1)
	return ret0, ret1
}
func (c *withTracing) Update(arg0 context.Context, arg1 *github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, arg2 k8s_io_apimachinery_pkg_apis_meta_v1.UpdateOptions) (*github_com_kyverno_kyverno_api_kyverno_v1alpha2.AdmissionReport, error) {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "Update"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("Update"),
	)
	defer span.End()
	arg0 = ctx
	ret0, ret1 := c.inner.Update(arg0, arg1, arg2)
	tracing.SetSpanStatus(span, ret1)
	return ret0, ret1
}
func (c *withTracing) Watch(arg0 context.Context, arg1 k8s_io_apimachinery_pkg_apis_meta_v1.ListOptions) (k8s_io_apimachinery_pkg_watch.Interface, error) {
	ctx, span := tracing.StartSpan(
		arg0,
		"",
		fmt.Sprintf("KUBE %s/%s/%s", c.client, c.kind, "Watch"),
		tracing.KubeClientGroupKey.String(c.client),
		tracing.KubeClientKindKey.String(c.kind),
		tracing.KubeClientOperationKey.String("Watch"),
	)
	defer span.End()
	arg0 = ctx
	ret0, ret1 := c.inner.Watch(arg0, arg1)
	tracing.SetSpanStatus(span, ret1)
	return ret0, ret1
}