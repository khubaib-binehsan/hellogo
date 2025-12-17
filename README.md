# hellogo

## Directories

- `cmd/`: Allows you to create main.go files that can then be used to build additional/specific binaries.

- `internal/`: Any package defined in internal is restricted by the compiler to only be available within the package, and cannot be used externally.

- `bar/`: Any arbitrary named file/folder that defines sub-package exportable to external users. Some use the convention `pkg/bar/` as well, but that's to one's own liking. Depending on the nature of project, pkg might give you better ergonomics as your directory is cleaner, but again, that is completely subjective.
