# BUILD

load("@rules_go//go:def.bzl", "go_library", "go_binary")
load("@gazelle//:def.bzl", "gazelle")

# Gazelle target to update BUILD files
gazelle(name = "gazelle")

# Define the library target for the internal package
go_library(
    name = "internal_pkg",
    srcs = ["internal_deps/internal_pkg.go"],
    importpath = "github.com/sujay-biradar-25/bazel_go_hello_world/internal_deps",
    visibility = ["//:__subpackages__"],
    deps = [
        "@com_github_bwmarrin_snowflake//:snowflake",
    ],
)

# Define the library target for the main library
go_library(
    name = "my_lib",
    srcs = ["libs/lib.go"],
    importpath = "github.com/sujay-biradar-25/bazel_go_hello_world/libs",
    visibility = ["//visibility:public"],
    deps = [
        ":internal_pkg",
    ],
)

# Define the binary target that depends on the main library
go_binary(
    name = "my_app",
    srcs = ["main.go"],
    importpath = "github.com/sujay-biradar-25/bazel_go_hello_world",
    visibility = ["//visibility:public"],
    deps = [":my_lib"],
)