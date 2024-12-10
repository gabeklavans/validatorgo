# Contributing to validatorgo

Welcome to validatorgo repository!! We appreciate your interest in contributing to this open source library and for helping our community grow.

## How to Contribute

### Code Contribution

In general, we follow the "fork-and-pull" Git workflow.

1. [Fork](https://docs.github.com/en/get-started/exploring-projects-on-github/contributing-to-a-project) the repository on GitHub
2. Clone the forked project to your local machine

```bash
git clone https://github.com/<your-username>/<repo-name>.git
cd <repo-name>
```

3. Work on your fork

   - Make your changes and additions
     - Most of your changes should be focused on root directory, the sanitizer package and the [README.md](https://github.com/bube054/validatorgo/blob/main/README.md).
   - Change or add tests if needed
     - Add or update tests in the relevant `*_test.go` files.
   - Run tests and make sure they pass

     ```bash
       go test ./...
     ```

   * Update documentation if necessary
     - Update [README.md](https://github.com/bube054/validatorgo/blob/main/README.md) and other go doc comments to reflect your changes.

4. Commit changes to your own branch

- Create a new branch for your work:
  ```bash
  git checkout -b <feature-or-fix-name>
  ```
- Stage and commit your changes
  ```bash
  git add .
  git commit -m "Describe your changes"
  ```

5. Merge the latest from "upstream" and resolve conflicts if any

- Add the upstream repository
  ```bash
  git remote add upstream https://github.com/<original-repo-owner>/<repo-name>.git
  ```
- Fetch and merge the latest changes
  ```bash
  git fetch upstream
  git merge upstream/main
  ```

6. Repeat step 3

- Re-test your changes after merging upstream updates to ensure everything still works as intended.

7. Push your work back up to your fork

```bash
git push origin <branch-name>
```

8. Submit a Pull Request (PR)

- Go to the original repository on GitHub and navigate to the Pull Requests tab.
- Click on "New Pull Request" and select the `dev` branch.
- Add a meaningful title and description for your changes.
- Submit the Pull Request for review.