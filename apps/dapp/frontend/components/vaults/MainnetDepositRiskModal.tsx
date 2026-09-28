"use client";

import React, { useState } from "react";
import { MAINNET_RISK_DISCLOSURES, LEGAL_TERMS_SUMMARY, MAINNET_TERMS_VERSION } from "@/lib/terms-and-disclosures";

interface MainnetDepositRiskModalProps {
  isOpen: boolean;
  onAccept: () => void;
  onCancel: () => void;
  vaultName?: string;
}

export function MainnetDepositRiskModal({
  isOpen,
  onAccept,
  onCancel,
  vaultName,
}: MainnetDepositRiskModalProps) {
  const [hasScrolledToBottom, setHasScrolledToBottom] = useState(false);
  const [agreed, setAgreed] = useState(false);

  if (!isOpen) return null;

  const handleScroll = (e: React.UIEvent<HTMLDivElement>) => {
    const target = e.currentTarget;
    const bottom = Math.abs(target.scrollHeight - target.scrollTop - target.clientHeight) < 5;
    if (bottom) {
      setHasScrolledToBottom(true);
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/70 backdrop-blur-sm p-4">
      <div className="bg-slate-900 border border-slate-800 rounded-2xl max-w-lg w-full p-6 text-slate-100 shadow-2xl flex flex-col max-h-[90vh]">
        <div className="flex items-center justify-between pb-4 border-b border-slate-800">
          <div>
            <h2 className="text-xl font-bold text-white">Mainnet Deposit Risk Disclosures</h2>
            <p className="text-xs text-slate-400 mt-0.5">Terms of Service & Legal Review (v{MAINNET_TERMS_VERSION})</p>
          </div>
        </div>

        {vaultName && (
          <div className="my-3 p-3 bg-indigo-950/40 border border-indigo-800/50 rounded-lg text-sm text-indigo-200">
            Depositing into: <span className="font-semibold text-white">{vaultName}</span>
          </div>
        )}

        <div
          onScroll={handleScroll}
          className="my-4 overflow-y-auto pr-2 space-y-4 text-sm text-slate-300 flex-1 border border-slate-800/80 rounded-xl p-4 bg-slate-950/50"
        >
          <div className="bg-amber-950/30 border border-amber-600/40 rounded-lg p-3 text-amber-200 text-xs">
            <strong>Legal Notice:</strong> Please review all risk disclosures carefully before committing real funds to Nester mainnet. Scroll to the bottom of the disclosures to enable confirmation.
          </div>

          <div className="whitespace-pre-line text-xs text-slate-400 leading-relaxed">
            {LEGAL_TERMS_SUMMARY.trim()}
          </div>

          <div className="space-y-3 pt-2">
            {MAINNET_RISK_DISCLOSURES.map((item, idx) => (
              <div key={item.id} className="p-3 bg-slate-900/80 border border-slate-800 rounded-lg">
                <h3 className="font-semibold text-white text-xs uppercase tracking-wide mb-1">
                  {idx + 1}. {item.title}
                </h3>
                <p className="text-xs text-slate-300 leading-relaxed">
                  {item.description}
                </p>
              </div>
            ))}
          </div>

          <div className="text-center text-xs text-slate-500 pt-2">
            End of Legal & Risk Disclosure Statement
          </div>
        </div>

        <div className="space-y-4 pt-2 border-t border-slate-800">
          <label className="flex items-start gap-3 cursor-pointer select-none">
            <input
              type="checkbox"
              checked={agreed}
              onChange={(e) => setAgreed(e.target.checked)}
              className="mt-0.5 h-4 w-4 rounded border-slate-700 bg-slate-800 text-indigo-600 focus:ring-indigo-500 focus:ring-offset-slate-900"
            />
            <span className="text-xs text-slate-300">
              I have read, reviewed, and agree to the Nester Terms of Service and acknowledge all smart contract risks, yield variability, and the absence of FDIC insurance.
            </span>
          </label>

          <div className="flex items-center justify-end gap-3">
            <button
              type="button"
              onClick={onCancel}
              className="px-4 py-2 rounded-xl text-xs font-medium text-slate-400 hover:text-white bg-slate-800 hover:bg-slate-700 transition"
            >
              Cancel
            </button>
            <button
              type="button"
              disabled={!agreed}
              onClick={onAccept}
              className="px-5 py-2 rounded-xl text-xs font-semibold text-white bg-indigo-600 hover:bg-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed transition shadow-lg shadow-indigo-600/20"
            >
              Proceed with Deposit
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
