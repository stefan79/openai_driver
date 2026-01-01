# Competition Analysis Specification

## Purpose
- Produce niche-specific competition analysis that compares each company's market position against identified competitors.
- Deliver actionable guidance on how the main company can strengthen or adjust its fit in its current niche, emphasizing its advantages while addressing gaps.

## Inputs
- **Market analysis per company**: Existing analyses scoped to the company's active niche (not aggregated across multiple niches).
- **Competitor links**: UVA references that map which companies compete with one another.
- **Niche definitions**: Canonical list of niches per company (source of truth for scoping reports).
- **Recent performance signals** (optional): Metrics such as growth rate, adoption, retention, pricing posture, and feature coverage to ground comparisons.

## Core Entities
- **Company Profile**: `{name, niche, positioning, pricing model, target segments, channels, differentiators, weaknesses, success metrics}`.
- **Competitor Set**: Derived from UVA references for the company’s niche.
- **Market Analysis**: Narrative plus structured findings for a single company in a single niche.
- **Comparison Report**: Generated artifact contrasting the main company with each competitor in the same niche.

## Workflow
1. **Ingest & Normalize**
   - Load market analyses per company; enforce one niche per analysis.
   - Resolve competitors via UVA references; filter to those active in the same niche.
   - Normalize terminology (niche names, segment labels, pricing terms).
2. **Niche-Scoped Refresh**
   - Trim or rewrite any multi-niche analysis into niche-specific coverage only.
   - Re-weight narrative to highlight the main company’s strengths within the focal niche while preserving objective risks.
3. **Comparison Construction**
   - For each competitor in the set, align on dimensions:
     - Positioning & target segments
     - Offering/feature set
     - GTM channels & partnerships
     - Pricing & packaging
     - Geographies or vertical focus
     - Operational scale and traction (if available)
   - Extract pros/cons per competitor relative to the main company.
   - Identify gaps (capability, pricing, distribution, messaging, geographic coverage).
4. **Recommendation Generation**
   - For each gap, propose adjustments: product changes, packaging tweaks, GTM moves, messaging pivots, or partnership options.
   - Flag quick wins vs. strategic bets; indicate required investment level.
5. **Report Assembly**
   - Build a per-competitor comparison report containing:
     - Niche overview (brief, scoped to the niche only)
     - Head-to-head summary table
     - Pros/cons list (competitor vs. main company)
     - Gaps and suggested closures
     - Adaptation guidance for market fit
   - Add an aggregated summary that rolls up common themes across competitors within the niche (no cross-niche aggregation).
6. **Quality & Bias Controls**
   - Ensure bias toward the main company is explicit and documented (e.g., “favor main company’s current niche strategy unless strong contradictory evidence”).
   - Maintain traceability: cite UVA references and source analyses for each claim.
   - Review for consistency of niche scoping and avoidance of cross-niche leakage.

## Outputs
- **Per-competitor comparison reports** (niche-scoped).
- **Niche summary** highlighting shared threats/opportunities and main-company advantages.
- **Actionable recommendation list** prioritized by effort/impact.
- **Audit trail** of UVA references and market-analysis sources used.

## Review Checklist
- Niche scope is singular and consistent across all sections.
- Competitor set matches UVA references for the niche.
- Pros/cons and gaps are evidence-backed and traceable.
- Recommendations are actionable, prioritized, and tied to identified gaps.
- Bias toward the main company is present but does not omit material risks.
- Terminology is normalized and comparable across competitors.

## Improvement Ideas
- Automate niche normalization and competitor mapping directly from UVA references.
- Introduce scoring/ranking per dimension (e.g., 1–5) to make deltas explicit.
- Add confidence levels and data freshness indicators per claim.
- Template the head-to-head summary for consistent readability across reports.
- Incorporate automated alerts when new competitors enter the niche or when a company expands into a new niche.

## Open Questions
- What is the canonical source for niche definitions, and how are conflicts resolved?
- Which metrics are available and reliable enough to quantify pros/cons (e.g., ARR bands, retention, adoption by segment)?
- Should recommendations include cost estimates or capacity constraints?
- How frequently should reports be regenerated when UVA references or market analyses change?
- Are there niches requiring custom dimensions beyond the standard comparison set?
