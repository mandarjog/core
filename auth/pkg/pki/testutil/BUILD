load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["util.go"],
    importpath = "istio.io/core/auth/pkg/pki/testutil",
    visibility = ["//visibility:public"],
    deps = ["@io_istio_auth//pkg/pki:go_default_library"],
)
