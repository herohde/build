// Package build contains utilities for build-time information.
package build

import "fmt"

var (
	// GitTree is to be defined by a linker argument as the result of
	// "git rev-list HEAD | wc -l | tr -d '[[:space:]]'".
	GitTree = "0"
	// GitHash is to be defined by a linker argument as the result of
	// "git rev-parse --short HEAD".
	GitHash = "nohash"
)

// Version holds the full build version, such as "1.4.2.120.4f54bc3d".
type Version struct {
	// Major is the major version.
	Major int
	// Minor is the minor version.
	Minor int
	// Micro is the Micro version.
	Micro int
	// Tree is size of the git commit tree, if defined. It is expected to be
	// a monotonically-increasing integer and be unique within a branch.
	Tree string
	// Hash is the short git commit hash, if defined, to avoid version
	// collisions for branched or privately built binaries.
	Hash string
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.%s.%s", v.Major, v.Minor, v.Micro, v.Tree, v.Hash)
}

// New creates a build version using the global GitTree and GitHash variables,
// which are expected to be defined externally as linker arguments.
func New(major, minor, micro int) Version {
	return Version{major, minor, micro, GitTree, GitHash}
}
