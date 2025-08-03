---
name: Bug Fix
about: Submit a PR that fixes a bug
title: 'fix: '
labels: ['bug', 'fix']
assignees: ''
---

# Conventional Commits Format
<!-- This PR title should follow the Conventional Commits specification: -->
<!-- https://www.conventionalcommits.org/en/v1.0.0/ -->
<!-- Format: fix(<optional scope>): <description> -->
<!-- Examples: -->
<!--   fix(login): handle invalid credentials properly -->
<!--   fix(api): prevent null pointer exception when response is empty -->

## Scope (optional)
<!-- What is the scope of this bug fix? (e.g. component or file name) -->

## Bug Description
<!-- Provide a clear and concise description of the bug being fixed -->

## Related Issue
<!-- Please link to the bug report issue this PR addresses -->
<!-- Format: Fixes #(issue) -->

## Root Cause
<!-- Describe the root cause of the bug -->

## Solution
<!-- Describe your solution to the bug -->
<!-- Explain why this is the best approach -->

## Before & After
<!-- If applicable, add screenshots or code snippets showing the before and after state -->

## Testing Performed
<!-- Describe the tests you've added or run to verify your fix -->
<!-- Include relevant details for manual testing if applicable -->
- [ ] Added unit tests that reproduce the bug
- [ ] Added integration tests (if applicable)
- [ ] Manually verified the fix with the following steps:
  <!-- List your verification steps -->

## Breaking Changes?
<!-- Will this change cause existing functionality to not work as expected? -->
<!-- If yes, please describe the impact and migration path for existing applications -->
- [ ] Yes (please describe below)
- [ ] No

## Potential Risks
<!-- Describe any potential risks that this change might introduce -->
<!-- If there are no risks, state "No known risks" -->

## Checklist
- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have made corresponding changes to the documentation (if applicable)
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published in downstream modules
- [ ] My PR title follows the Conventional Commits format

## Additional Notes
<!-- Add any other context about the PR here -->
