# Agent guidance

## Generated documentation
- The root readme.md is generated and says not to edit it manually. Determine the authoritative template/generator (readme.md.tmpl or project tooling), change that source, then regenerate and verify when documentation needs changing.
- Do not commit regenerated unrelated artifacts unless the diff is understood and necessary.

## Image/frame invariants
- Preserve the documented image.Image and draw.Draw-compatible behaviour and nine-slice frame layout: corners remain stable, while edges and centre are repeated or stretched according to the chosen implementation.
- Test non-zero rectangle origins, small/degenerate rectangles, alpha composition, clipping and bounds with small in-memory images and explicit pixel assertions. Golden PNGs are appropriate only for actual whole-image behaviour.
- Avoid nondeterministic imagery and incidental real files in pure layout tests. Keep existing public API call sites compatible.

## Agent handoff
Inspect code, README template, any generation commands and tests, then attach validated gaps to issue #66. Separate implementation changes from this rules-only PR.
## Engineering reference and follow-up

- Adopt only relevant guidance from https://arran4.github.io/blog/post/2026/055-suggested-coding-standards-for-repositories-and-agents/ and https://arran4.github.io/blog/post/2026/020-optional-dependency-injection-via-type-switched-variadic-args/ .
- This documentation is not evidence of repository-wide compliance. A coding-agent pass must report checked files, observed gaps, tests/CI and existing issue overlap to https://github.com/arran4/golang-frame/issues/66 before any broad implementation.
- Do not merge implementation automatically; request review and follow repository-specific instructions.
