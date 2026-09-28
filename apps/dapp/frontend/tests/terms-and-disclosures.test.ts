import { describe, it, expect } from "vitest";
import { MAINNET_RISK_DISCLOSURES, MAINNET_TERMS_VERSION } from "@/lib/terms-and-disclosures";

describe("Mainnet Terms and Risk Disclosures", () => {
  it("defines a valid terms version", () => {
    expect(MAINNET_TERMS_VERSION).toBeDefined();
    expect(typeof MAINNET_TERMS_VERSION).toBe("string");
  });

coreCheck: {
    it("includes essential risk disclosures for mainnet deposit", () => {
      expect(MAINNET_RISK_DISCLOSURES.length).toBeGreaterThanOrEqual(3);
      
      const ids = MAINNET_RISK_DISCLOSURES.map((d) => d.id);
      expect(ids).toContain("smart-contract-risk");
      expect(ids).toContain("no-fdic-insurance");
      expect(ids).toContain("yield-variability");
    });
  }
});
