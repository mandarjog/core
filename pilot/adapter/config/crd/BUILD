load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "config.go",
        "controller.go",
        "conversion.go",
        "types.go",
    ],
    importpath = "istio.io/core/pilot/adapter/config/crd",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@com_github_hashicorp_go_multierror//:go_default_library",
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//platform/kube:go_default_library",
        "@io_k8s_apiextensions_apiserver//pkg/apis/apiextensions/v1beta1:go_default_library",
        "@io_k8s_apiextensions_apiserver//pkg/client/clientset/clientset:go_default_library",
        "@io_k8s_apimachinery//pkg/api/errors:go_default_library",
        "@io_k8s_apimachinery//pkg/apis/meta/v1:go_default_library",
        "@io_k8s_apimachinery//pkg/runtime:go_default_library",
        "@io_k8s_apimachinery//pkg/runtime/schema:go_default_library",
        "@io_k8s_apimachinery//pkg/runtime/serializer:go_default_library",
        "@io_k8s_apimachinery//pkg/util/wait:go_default_library",
        "@io_k8s_apimachinery//pkg/util/yaml:go_default_library",
        "@io_k8s_apimachinery//pkg/watch:go_default_library",
        "@io_k8s_client_go//plugin/pkg/client/auth/gcp:go_default_library",
        "@io_k8s_client_go//plugin/pkg/client/auth/oidc:go_default_library",
        "@io_k8s_client_go//rest:go_default_library",
        "@io_k8s_client_go//tools/cache:go_default_library",
        "@io_k8s_client_go//tools/clientcmd:go_default_library",
    ],
)

genrule(
    name = "crd",
    srcs = ["config.go"],
    outs = ["types.go"],
    cmd = "adapter/config/crd/generate.sh > $@",
    tools = ["generate.sh"],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = [
        "client_test.go",
        "controller_test.go",
        "conversion_test.go",
    ],
    data = ["//platform/kube:kubeconfig"],
    importpath = "istio.io/core/pilot/adapter/config/crd",
    library = ":go_default_library",
    deps = [
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//platform/kube:go_default_library",
        "@io_istio_pilot//test/mock:go_default_library",
        "@io_istio_pilot//test/util:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    size = "small",
    srcs = ["config_test.go"],
    importpath = "istio.io/core/pilot/adapter/config/crd_test",
    deps = [
        "@io_istio_pilot//adapter/config/crd:go_default_library",
        "@io_k8s_apimachinery//pkg/apis/meta/v1:go_default_library",
    ],
)
