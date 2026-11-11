# Changelog

## 0.1.0-alpha.3

- Add number reuse, phone readiness, and SMS reception start for Activation V2.
- Preserve method-specific HTTP 202 responses, retryAfter and nullable order expiry.
- Cover all 61 operations while preserving the existing 58 contracts.

## 0.1.0-alpha.2

- Add all-country and country-specific compact Activation counts with optional provider ratings.
- Preserve raw counts, rating fields, country casing, documented authentication and all existing method names.
- Cover all 58 operations with local mock and wire-contract tests.

## 0.1.0-alpha.1

- First public alpha covering all 56 Activation V2, Rental and Private Email operations.
- Live reference methods, original response values and explicit transport errors.
- No automatic retries for paid or state-changing operations.
