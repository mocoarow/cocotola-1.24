import { createCookieSessionStorage, redirect } from "react-router";

const sessionSecret = process.env.SESSION_SECRET;
if (!sessionSecret) {
  throw Error("SESSION_SECRET must be set");
}

type LoginData = {
  state: string;
  nonce: string;
  codeVerifier: string;
};

type UserData = {
  accessToken: string;
  refreshToken: string;
  sub: string;
  loginId: string;
  name: string;
};

const loginDataSession = createCookieSessionStorage<LoginData>({
  cookie: {
    name: "loginData",
    httpOnly: true,
    maxAge: 60 * 10,
    path: "/",
    sameSite: "lax",
    secrets: [sessionSecret],
    secure: process.env.NODE_ENV === "production",
  },
});
async function getLoginDataSession(request: Request) {
  return await loginDataSession.getSession(request.headers.get("Cookie"));
}
const {
  commitSession: commitLoginDataSession,
  destroySession: destroyLoginDataSession,
} = loginDataSession;

const userDataSession = createCookieSessionStorage<UserData>({
  cookie: {
    name: "userData",
    httpOnly: true,
    maxAge: 60 * 10,
    path: "/",
    sameSite: "lax",
    secrets: [sessionSecret],
    secure: process.env.NODE_ENV === "production",
  },
});

const {
  commitSession: commitUserDataSession,
  destroySession: destroyUserDataSession,
} = userDataSession;

async function getUserDataSession(request: Request) {
  return await userDataSession.getSession(request.headers.get("Cookie"));
}

async function requireAuthUser(request: Request) {
  console.log("requireAuthUser");

  const sessionUserData = await getUserDataSession(request);

  if (sessionUserData) {
    const loginId = sessionUserData.get("loginId");
    const name = sessionUserData.get("name");
    if (loginId && name) {
      return {
        loginId,
        name,
      };
    }
  }

  const _sessionLoginData = await getLoginDataSession(request);

  // TODO: fetch userInfo with accessToken

  // TODO: refresh accessToken with refreshToken

  throw redirect("/login", {
    headers: {
      "Set-Cookie": await destroyUserDataSession(sessionUserData),
    },
  });
}

export {
  getLoginDataSession,
  commitLoginDataSession,
  destroyLoginDataSession,
  getUserDataSession,
  commitUserDataSession,
  destroyUserDataSession,
  requireAuthUser,
};
