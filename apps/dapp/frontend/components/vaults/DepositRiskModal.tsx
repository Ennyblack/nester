import React, { useState } from "react";

interface DepositRiskModalProps {
  isOpen: boolean;
  onAccept: () => void;
  onCancel: () => void;
}

export const DepositRiskModal: React.FC<DepositRiskModalProps> = ({
  isOpen,
  onAccept,
  onCancel,
}) => {
  const [confirmed, setConfirmed] = useState(false);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/70 p-4">
      <div className="max-w-lg w-full bg-slate-900 border border-slate-800 rounded-xl p-6 text-slate-100 shadow-xl">
        <h2 className="text-xl font-bold mb-4 text-amber-400">
          Mainnet Risk Disclosures & Terms
        </h2>
        <div className="space-y-3 text-sm text-slate-300 max-h-96 overflow-y-auto pr-2 mb-6">
          <p>
            Before depositing real funds into Nester Smart Vaults, please review
            and acknowledge the following critical disclosures:
          </p>
          <div className="bg-slate-800/60 p-3 rounded-lg border border-slate-700/50">
            <h3 className="font-semibold text-slate-200 mb-1">
              1. No FDIC or Government Insurance
            </h3>
            <p>
              Your deposits are <strong>not insured</strong> by the FDIC, SIPC,
              or any government agency. You bear the full risk of loss.
            </p>
          </div>
          <div className="bg-slate-800/60 p-3 rounded-lg border border-slate-700/50">
            <h3 className="font-semibold text-slate-200 mb-1">
              2. Smart Contract & DeFi Risk
            </h3>
            <p>
              Smart contracts and underlying yield adapters may contain bugs or
              be subject to exploits and third-party protocol failures.
            </p>
          </div>
          <div className="bg-slate-800/60 p-3 rounded-lg border border-slate-700/50">
            <h3 className="font-semibold text-slate-200 mb-1">
              3. Yield Variability
            </h3>
            <p>
              Target APYs and yields are variable, depend on market conditions,
              and are never guaranteed.
            </p>
          </div>
        </div>
        <div className="flex items-center space-x-3 mb-6">
          <input
            type="checkbox"
            id="risk-checkbox"
            checked={confirmed}
            onChange={(e) => setConfirmed(e.target.checked)}
            className="w-4 h-4 accent-amber-500 rounded"
          />
          <label htmlFor="risk-checkbox" className="text-xs text-slate-300 select-none">
            I have read, understood, and legally accept the Nester Terms of
            Service and Risk Disclosures.
          </label>
        </div>
        <div className="flex justify-end space-x-3">
          <button
            type="button"
            onClick={onCancel}
            className="px-4 py-2 text-sm bg-slate-800 hover:bg-slate-700 rounded-lg text-slate-300 transition-colors"
          >
            Cancel
          </button>
          <button
            type="button"
            disabled={!confirmed}
            onClick={onAccept}
            className="px-4 py-2 text-sm bg-amber-500 hover:bg-amber-600 disabled:opacity-50 disabled:cursor-not-allowed rounded-lg font-medium text-slate-950 transition-colors"
          >
            Proceed to Deposit
          </button>
        </div>
      </div>
    </div>
  );
};
