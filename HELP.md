# beacon-model help

## Retaging
1. list all remote tags
git ls-remote --tags

2. delete local tag
git tag -d v0.0.1

3. push tag deletion to remote
git push origin :refs/tags/v0.0.1

4. tag local branch again
git tag v0.0.1

5. push tag to remote
git push origin tag v0.0.1
