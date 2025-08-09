import * as client from "openid-client";

const googleClientId = process.env.GOOGLE_CLIENT_ID || "";
if (!googleClientId) {
  throw new Error("GOOGLE_CLIENT_ID must be set");
}
const googleClientSecret = process.env.GOOGLE_CLIENT_SECRET || "";
if (!googleClientSecret) {
  throw new Error("GOOGLE_CLIENT_SECRET must be set");
}

export async function getAuthConfig(): Promise<client.Configuration> {
  const server = new URL(
    "https://accounts.google.com/.well-known/openid-configuration",
  );

  return await client.discovery(server, googleClientId, googleClientSecret);
}
