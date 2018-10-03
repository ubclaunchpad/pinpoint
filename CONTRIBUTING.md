# :books: Contributing

This document outlines key considerations and tips for anyone contributing to
Pinpoint.

## :hammer_and_pick: Development

Development instructions are available in the repository
[README](https://github.com/ubclaunchpad/pinpoint/blob/master/README.md).

## :robot: Submitting Changes

All changes should be made via
[pull requests](https://help.github.com/articles/about-pull-requests/), which
are reviewed by one or more members of the maintaining team before being
merged. In additional to reviews, all pull requests must pass our 
[continuous integration](https://en.wikipedia.org/wiki/Continuous_integration),
which consists of a number of tasks and tests run by
[Travis CI](https://travis-ci.com/ubclaunchpad/pinpoint).

Submitting a pull request consists of several steps:

### 1. Creating and Assigning an Issue

Check if an [Issue](https://github.com/ubclaunchpad/pinpoint/issues)
(also referred to as 'ticket') exists for the change you want to make - if there
isn't, create one with a description of the task. You should then 'Assign' the
ticket to yourself, so that everyone else is aware that this task is being
worked on.

Relevant issues and pull requests should be included as part a
[project](https://github.com/ubclaunchpad/pinpoint/projects), which provides us
with a kanban board with an overview of the project's progress.

### 2. Creating a Branch

Pull requests are made to request permission to merge from one branch into
another, typically `master`. This applies the changes from the 'source' branch
onto the 'target' branch.

Before you create a branch, you want to make sure you are branching off at
the latest commit, so that you are not missing any changes that might affect
your work. Navigate to where the Pinpoint source code is and run the following
commands:

```sh
$> git checkout master
$> git pull
```

Now that your codebase is up to date, you can create a branch. Branches should
be named as follows:

```
<component>/<description>
```

`<component>` refers to the part of the codebase being worked on, such as
`core`, `gateway`, `frontend`, `client`, etc. `<description>` should be a very
brief title that describes the changes you want to include in this branch and
pull request. For example, to add a login page to the frontend, your branch
name might be called `frontend/login-page`.

To create and switch to this new branch:

```sh
$> git checkout -b <my-new-branch-name>
```

You can now make changes and commit them to this branch, without fear of
interfering with other people's work. To create a commit:

```sh
$> git add .     # 'stage' all changes
$> git commit -m "<your commit message here>"
```

Commit messages are very important - 
[this guide](https://chris.beams.io/posts/git-commit/#seven-rules) provides
a great overview of what makes a good commit message.

#### 3. Submit a Pull Request

To submit a pull request, you must first push you changes onto GitHub:

```sh
$> git push origin HEAD
```

Once you have done this, you can visit 
[the pull request tab](https://github.com/ubclaunchpad/pinpoint/pulls) of the
repository to create a pull request. Simply select your branch as the `compare`
branch. There is a pull request template available, where you should fill in
the blanks to describe the changes you have made.

There is a more detailed guide available in the official
[UBC Launch Pad documentation](https://github.com/ubclaunchpad/pinpoint/pulls) -
feel free to refer to that as well!
