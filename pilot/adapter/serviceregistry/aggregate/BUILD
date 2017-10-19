load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["controller.go"],
    importpath = "istio.io/core/pilot/adapter/serviceregistry/aggregate",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//platform:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = ["controller_test.go"],
    importpath = "istio.io/core/pilot/adapter/serviceregistry/aggregate",
    library = ":go_default_library",
    deps = [
        "@io_istio_pilot//model:go_default_library",
        "@io_istio_pilot//platform:go_default_library",
        "@io_istio_pilot//test/mock:go_default_library",
    ],
)
