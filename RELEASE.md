# Steps to Release dgo

## Releasing a Minor Version

1. Checkout the commit that you want to tag with the new release.

2. Run the following command to create an annotated tag:

   ```sh
   git tag <new tag>
   ```

3. Push the tag to GitHub:

   ```sh
   git push origin <new tag>
   ```

## Releasing a Major Version

1. Update the `go.mod` file to the module name with the correct version.

2. Change all the import paths to import `v<new major version>`. For example, if the current import
   path is `"github.com/dgraph-io/dgo/v200"`. When we release v201.07.0, we would replace the import
   paths to `"github.com/dgraph-io/dgo/v201"`.

3. Update [Supported Version](https://github.com/hypermodeinc/dgo/#supported-versions).

4. Commit all the changes and get them merged to master branch.

   Now, follow the [Releasing a Minor Version](#releasing-a-minor-version) as above.

   Note that, now you may have to also change the import paths in the applications that use dgo
   including dgraph and raise appropriate PR for them.
