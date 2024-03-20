package constants

const FORGOT_MAIL string = `Subject: Urgent Action Required: Secure Your Expense Monster Account

Hi [USER],

We detected a request to regain access to your Expense Monster account. As security is our top priority, we've taken precautions to verify your identity.

This email contains a one-time pass-code (OTP) that will grant you temporary access to reset your password. Remember, this code is highly confidential and serves as the key to regaining control of your account.

Here's your OTP: [OTP]

Please note: This code will expire in [MIN] minutes for your security. You have a maximum of [ATTEMPT] attempts to enter the correct OTP before your account is temporarily locked for your protection.

If you did not initiate this request, do not proceed further.  For your account's safety, we recommend changing your password immediately upon regaining access.

Act swiftly to reclaim your account.

Sincerely,
Team Expense Monster`

const MAX_ATTEMPT_FORGOT int = 3
const MAX_TIME_FORGOT_MINS int = 5
