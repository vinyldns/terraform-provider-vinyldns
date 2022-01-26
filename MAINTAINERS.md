# Maintainers

The current maintainers for this project are:

- Ryan Emerle ([@remerle](https://github.com/remerle))
- Sriram Ramakrishnan ([@sramakr](https://github.com/sramakr))
- Jim Wakemen ([@jwakemen](https://github.com/jwakemen))

## Releases

Releases are controlled by
a [GitHub Action](https://github.com/vinyldns/terraform-provider-vinyldns/blob/master/.github/workflows/release.yml). To
trigger a new release, you simply need to create a new tag.

```shell
git tag v0.11.0
git push origin --tags
```

Tags should be prefixed with `v` and following [semantic versioning](https://semver.org/) standards for the `major`
, `minor`, and `patch` revision numbers.