load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["secret.go"],
    importpath = "istio.io/core/auth/pkg/pki/ca/controller",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@io_istio_auth//pkg/pki:go_default_library",
        "@io_istio_auth//pkg/pki/ca:go_default_library",
        "@io_k8s_apimachinery//pkg/api/errors:go_default_library",
        "@io_k8s_apimachinery//pkg/apis/meta/v1:go_default_library",
        "@io_k8s_apimachinery//pkg/fields:go_default_library",
        "@io_k8s_apimachinery//pkg/runtime:go_default_library",
        "@io_k8s_apimachinery//pkg/watch:go_default_library",
        "@io_k8s_client_go//kubernetes/typed/core/v1:go_default_library",
        "@io_k8s_client_go//pkg/api/v1:go_default_library",
        "@io_k8s_client_go//tools/cache:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = ["secret_test.go"],
    importpath = "istio.io/core/auth/pkg/pki/ca/controller",
    library = ":go_default_library",
    deps = [
        "@io_istio_auth//pkg/pki/ca:go_default_library",
        "@io_k8s_apimachinery//pkg/apis/meta/v1:go_default_library",
        "@io_k8s_apimachinery//pkg/runtime/schema:go_default_library",
        "@io_k8s_client_go//kubernetes/fake:go_default_library",
        "@io_k8s_client_go//pkg/api/v1:go_default_library",
        "@io_k8s_client_go//testing:go_default_library",
    ],
)
