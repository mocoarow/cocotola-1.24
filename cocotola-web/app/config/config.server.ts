export const serverConfigRedirectUri = process.env.REDIRECT_URI||"";
if (!serverConfigRedirectUri) {
  throw new Error("REDIRECT_URI is not set");
}

