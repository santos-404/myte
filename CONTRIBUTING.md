# Contribution guidelines

Everyone is welcome to contribute; whether it's ideas, bug reports, code, or feedback. However, to ensure quality and consistency, please follow the steps below before submitting any contribution.

## Before you start

* Be kind and respectful.
* Understand that this is not a stable or production-ready project — breaking changes may occur.
* All tests **must pass** before submitting your pull request.

## How to contribute

1. **Fork** the repository to your GitHub account.

2. **Open an issue** describing the bug or feature you'd like to work on.

3. **Wait for confirmation or discussion**, especially if you're proposing a new feature. You can start with the following steps without waiting, but then you must know that not every change must be approved.

4. Once approved, **create a new branch** from the latest `main`.

5. **Reference the issue** in your commits (e.g. `add X functionality (#42)`) and in your pull request description.

6. Ensure all tests pass by running:

   ```bash
   go test -v ./...
   ```

   > All tests must pass. Your PR will not be accepted otherwise.

7. **Do not modify existing test files**. These are locked to ensure consistent behavior and compatibility.

   If your contribution requires **adding new tests**, include them in a **new `_test.go` file**, and justify their inclusion in the PR description. Avoid modifying or duplicating tests from the existing suite.

8. Push your branch and **open a pull request** targeting `main`.

## Automated workflow

A GitHub Actions workflow will automatically run the test suite on every pull request. If any test fails, you’ll need to fix it before the PR can be merged.

---

We appreciate your effort in making this project better. Thank you for contributing!
