export interface RiskDisclosureItem {
  id: string;
  title: string;
  description: string;
}

export const MAINNET_TERMS_VERSION = "1.0.0";

export const MAINNET_RISK_DISCLOSURES: RiskDisclosureItem[] = [
  {
    id: "smart-contract-risk",
    title: "Smart Contract & Protocol Risk",
    description: "Funds deposited into Nester Smart Vaults interact with on-chain Soroban smart contracts and underlying DeFi lending/staking protocols. Smart contracts may contain undiscovered vulnerabilities, bugs, or exploits that could result in the permanent loss of deposited assets.",
  },
  {
    id: "no-fdic-insurance",
    title: "No FDIC or Government Insurance",
    description: "Deposits with Nester are self-custodial and decentralized on the Stellar network. They are NOT bank deposits, are NOT insured by the Federal Deposit Insurance Corporation (FDIC), the Securities Investor Protection Corporation (SIPC), or any other government agency, and carry no sovereign or institutional backstops.",
  },
  {
    id: "yield-variability",
    title: "Yield Variability & Market Risk",
    description: "Target APYs and yields are variable and fluctuate in real time based on on-chain supply, demand, and liquidity conditions across decentralized financial markets. Past performance and projected returns do not guarantee future yields, and principal values may experience impermanent loss or negative returns.",
  },
  {
    id: "self-custody-responsibility",
    title: "Self-Custody & User Responsibility",
    description: "You maintain full, exclusive control and custody of your private keys and self-custodial wallet. Nester cannot recover lost keys, reverse transactions, or restore funds sent to incorrect addresses. You are solely responsible for all transaction fees and actions taken from your wallet.",
  },
];

export const LEGAL_TERMS_SUMMARY = `
By accessing or depositing real funds into Nester mainnet vaults, you agree to the Nester Terms of Service and acknowledge that you have read, understood, and accepted all associated risk disclosures. You represent that you are legally permitted to use decentralized financial applications in your jurisdiction.
`;
