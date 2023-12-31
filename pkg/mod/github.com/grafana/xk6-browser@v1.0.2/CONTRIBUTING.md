# Contributing to xk6-browser

Thank you for your interest in contributing to xk6-browser!

(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧

Before you begin, make sure to familiarize yourself with the [Code of Conduct](CODE_OF_CONDUCT.md). If you've previously contributed to other open source project, you may recognize it as the classic [Contributor Covenant](https://contributor-covenant.org/).

If you want to chat with the team or the community, you can [join our community forums](https://community.k6.io/c/xk6-browser/14).

> **Note:** To disclose security issues, refer to [SECURITY.md](SECURITY.md).

## Filing issues

Don't be afraid to file issues! Nobody can fix a bug we don't know exists, or add a feature we didn't think of.

The worst that can happen is that someone closes it and points you in the right direction.

That said, "how do I..."-type questions are often more suited for community forums.

## Contributing code

If you'd like to contribute code to xk6-browser, this is the basic procedure. Make sure to follow the [style guide](#code-style) described below.

1. Find an issue you'd like to fix. If there is none already, or you'd like to add a feature, please open one, and we can talk about how to do it.  Out of respect for your time, please start a discussion regarding any bigger contributions either in a GitHub Issue, in the community forums **before** you get started on the implementation.
  
   Remember, there's more to software development than code; if it's not properly planned, stuff gets messy real fast.

2. Create a fork and open a feature branch - `feature/my-cool-feature` is the classic way to name these, but it really doesn't matter.

3. Create a pull request!

4. Sign the [Contributor License Agreement](https://cla-assistant.io/grafana/xk6-browser) (the process is integrated with the pull request flow through cla-assistant.io).

5. We will discuss implementation details until everyone is happy, then a maintainer will merge it.

### Development setup

To get a basic development environment for Go and xk6-browser up and running, first make sure you have **[Git](https://git-scm.com/downloads)** and **[Go](https://golang.org/doc/install)** (see our [go.mod](https://github.com/grafana/xk6-browser/blob/master/go.mod#L3) for minimum required version) installed and working properly.

We recommend using the Git command-line interface to download the source code for the xk6-browser repo:

* Open a terminal and run `git clone https://github.com/grafana/xk6-browser.git`. This command downloads xk6-browser's source code to a new `xk6-browser` directory in your current directory.
* Open the `xk6-browser` directory in your favorite code editor.

For alternative ways of cloning the xk6-browser repository, please refer to [GitHub's cloning a repository](https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/cloning-a-repository) documentation.

#### Running the linter

We make use of the [golangci-lint](https://github.com/golangci/golangci-lint) tool to lint the code in CI. The actual version you can find in our [`.golangci.yml`](https://github.com/grafana/xk6-browser/blob/master/.golangci.yml#L1). To run it locally, first [install it](https://golangci-lint.run/usage/install/#local-installation), then run:

```bash
make lint
```

You can also run the linter inside the docker container, which will benefit from the version of the linter being the same as it will be in CI.

```bash
make ci-like-lint
```

#### Running the test suite

To exercise the entire test suite, please run the following command

```bash
make tests
```

#### Dependencies

It's best to avoid adding or updating dependencies. If you think your PR must have a new or updated dependency we will take a look on a case by case basis, but we will likely ask that you find an alternate solution without needing to require a new or updated dependency.

It's worth noting that we need to abide by the [dependency](https://github.com/grafana/k6/blob/master/Dependencies.md) guidelines in the [k6](https://github.com/grafana/k6) codebase.

#### Code style

As you'd expect, please adhere to good ol' `gofmt` (there are plugins for most editors that can autocorrect this), but also `gofmt -s` (code simplification), and don't leave unused functions laying around.

Continuous integration will catch all of this if you don't, and it's fine to just fix linter complaints with another commit, but you can also run the linter yourself:

```bash
make check
```

Comments in the source should wrap at 100 characters, but there's no maximum length or need to be brief here - please include anything one might need to know in order to understand the code, that you could reasonably expect any reader to not already know (you probably don't need to explain what a goroutine is).

#### Commit format

Please read [this superb article from Github](https://github.blog/2022-06-30-write-better-commits-build-better-projects/) to get an idea of how we work with commits.

The following are the guidelines that we adhere to:

##### What to include in a commit?
- Keep it as small as possible.
- Keep it as related as possible.
- Do not mix refactoring and reformatting with other changes.
   - Move a function to another place in one commit.
   - Fix a bug in another commit.
   - Refactor it in yet another commit.

##### Commit subjects

- Use imperative language.
- Keep it to 50 characters max.
- Add a blank line if you need to explain more (please do 🙏).
- Capitalize the subject line.
- Do not end the subject line with a period.
- If it's a commit based on review comments then add the PR number e.g. Apply review comments (#400).
- Use the following commit verbs:
   - Add: Create a capability like a feature, test, dependency.
   - Remove: Remove a capability like a feature, test, dependency.
   - Fix: Fix an issue like a bug or typo.
   - Refactor: A code change that is just refactoring.
   - Reformat: Refactor of formatting, like a omit whitespace.
   - Rename: Rename something.
   - Optimize: Refactor of performance, like a speed up code.
   - Update: Update a dependency.
   - Bump: Increase the version.
   - Release: Before releasing a new version.

##### Commit description

- Explain what and why not how.
- Write your code so that it explains the how.
- Limit each line at 72 characters max.
- Finish your commit with the following Github messages if it does something with an issue:
   - Fixes: #N.
   - Closes: #N.
   - References (or See): #N.
   - Resolves: link to comment.

##### Closing an issue

If your commit closes an issue, please [close it with your commit message](https://help.github.com/articles/closing-issues-via-commit-messages/), for example:

```text
Added this really rad feature

Closes #420
```

#### Language and text formatting

Any human-readable text you add must be non-gendered and should be fairly concise without devolving into grammatical horrors, dropped words, and shorthands. This isn't Twitter, you don't have a character cap, but don't write a novel where a single sentence would suffice.

If you're writing a longer block of text to a terminal, wrap it at 80 characters - this ensures it will display properly at the de facto default terminal size of 80x25. As an example, this is the help text of the `scale` sub-command from k6:

```text
   Scale will change the number of active VUs of a running test.

   It is an error to scale a test beyond vus-max; this is because instantiating
   new VUs is a very expensive operation, which may skew test results if done
   during a running test. To raise vus-max, use --max/-m.

   Endpoint: /v1/status
```
