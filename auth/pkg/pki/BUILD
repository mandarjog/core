load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "crypto.go",
        "san.go",
    ],
    importpath = "istio.io/core/auth/pkg/pki",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "crypto_test.go",
        "san_test.go",
    ],
    importpath = "istio.io/core/auth/pkg/pki",
    library = ":go_default_library",
)
