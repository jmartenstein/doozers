# Red / Green / Refactor (TDD)

The Red/Green/Refactor cycle is a foundational framework in Test-Driven Development (TDD) that drives software design through short, iterative development cycles.

## The Three Phases

1. **Red (Think about *what* to develop):**
    - **Action:** Write a test for a specific feature or function before any implementation code exists.
    - **Goal:** Define the expected behavior. The test must fail (turning "red") because the feature hasn't been built yet. This failure confirms the test is valid and informs your implementation approach.

2. **Green (Think about *how* to make it work):**
    - **Action:** Write the minimum amount of implementation code required to make the test pass.
    - **Goal:** Get the test to pass (turning "green") as quickly as possible. At this stage, focus on functionality rather than optimization or "perfect" code.

3. **Refactor (Think about *how* to improve it):**
    - **Action:** Clean up and optimize both the implementation code and the test suite while ensuring the tests remain green.
    - **Goal:** Improve efficiency, readability, and maintainability. The existing test acts as a safety net, immediately alerting you if your changes break the functionality.

## Refactoring Best Practices

Before moving back to the "Red" phase, evaluate the following:
- **Clarity:** Can I make the implementation or test suite more descriptive or expressive?
- **DRY (Don't Repeat Yourself):** Can I reduce duplication in the code?
- **Efficiency:** Can I implement the logic more efficiently (e.g., faster algorithms)?
- **Isolation:** Are the tests isolated and providing reliable feedback?

## Benefits
- **Confidence:** Tests provide a safety net for making changes.
- **Early Bug Detection:** Catch bugs early and prevent regressions.
- **Robustness:** Results in a more complete and reliable software solution.
