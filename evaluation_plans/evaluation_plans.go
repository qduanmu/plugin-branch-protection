package evaluation_plans

import (
    "github.com/ossf/gemara/layer4"
    "github.com/MYORG/plugin-branch_protection/evaluation_plans/reusable_steps"
)

var (
    // Add more entries here if other catalogs or catalog versions are adopted by the plugin
    // Then use orchestrator.AddEvaluationSuite to make these available to the user
    SC_CODE = map[string][]layer4.AssessmentStep{
        // For best results, ensure every assessment id is represented in this map
        "SC-CODE-AC-01.01": {
            /* Repository owners MUST require pull requests or merge requests before merging
changes to protected branches. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-01.02": {
            /* Repository owners MUST enforce a minimum number of approvals (≥1) from reviewers
who are different from the author before allowing changes to be merged into
protected branches. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-01.03": {
            /* Repository owners MUST block force pushes to protected branches. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-01.04": {
            /* Repository owners MUST prevent bypassing of branch protection settings. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-01.05": {
            /* Repository owners MUST require CODEOWNER reviews for changes to protected branches. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-01.06": {
            /* Repository owners MUST restrict merge permissions to authorized roles (e.g., maintainers)
for protected branches. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-02.01": {
            /* Repository owners MUST enforce 2FA at the organization or group level for all
users to restrict repository access to authenticated users only. */
            reusable_steps.NotImplemented,
        },
        "SC-CODE-AC-02.02": {
            /* Repository owners MUST enforce 2FA at the instance level for all users to restrict
repository access to authenticated users only. */
            reusable_steps.NotImplemented,
        },
        
    }
)
