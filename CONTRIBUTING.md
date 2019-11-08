# Contributing to terraform-provider-vinyldns

The following is a set of guidelines for contributing. These are mostly guidelines, not rules. Use your best judgment and feel free to propose changes to this document in a pull request.

## Code of Conduct

This project and everyone participating in it is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior by submitting an issue under the github repo.

## Reporting Bugs

Bugs are tracked as [GitHub issues](https://guides.github.com/features/issues/). Please complete [the template](ISSUE_TEMPLATE.md) when creating an issue. This helps resolve issues faster.

> **Note:** If you find a **Closed** issue that seems like it is the same thing that you're experiencing, open a new issue and include a link to the original issue in the body of your new one.

Please also consider the following when creating reporting a bug:

* **Use a clear and descriptive title** for the issue to identify the problem.
* **Describe the exact steps which reproduce the problem** in as many details as possible. When listing steps, **don't just say what you did, but explain how you did it**. For example, if you moved the cursor to the end of a line, explain if you used the mouse, or a keyboard shortcut.
* **Provide specific examples to demonstrate the steps**. Include details such as a strack trace, links to files or GitHub projects, or copy/pasteable snippets that illustrate the problem. If you're providing snippets in the issue, use [Markdown code blocks](https://help.github.com/articles/markdown-basics/#multiple-lines).
* **Describe the behavior you observed after following the steps** and point out what exactly is the problem with that behavior.
* **Explain the behavior you expected to see instead and why.**

Provide more context by answering these questions:

* **Did the problem start happening recently** (e.g. after updating to a new version) or was this always a problem?
* If the problem started happening recently, **can you reproduce the problem in an older version?** What's the most recent version in which the problem doesn't happen?
* **Can you reliably reproduce the issue?** If not, provide details about how often the problem happens and under which conditions it normally happens.

## Pull Requests

* Complete the [the required template](PULL_REQUEST_TEMPLATE.md)
* Do not include issue numbers in the PR title

## Documentation
* Documentation lives in the `docs` directory and the site is deployed with GitHub pages
* To run the documentation site locally: `cd docs; python -m SimpleHTTPServer 3000`

### Contributor License Agreement

Before Comcast merges your code into the project you must sign the [Comcast Contributor License Agreement (CLA)](https://gist.github.com/ComcastOSS/a7b8933dd8e368535378cda25c92d19a).

If you haven't previously signed a Comcast CLA, you'll automatically be asked to when you open a pull request. Alternatively, we can send you a PDF that you can sign and scan back to us. Please create a new GitHub issue to request a PDF version of the CLA.

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line
