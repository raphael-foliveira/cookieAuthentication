import { NextRequest, NextResponse } from "next/server";

export function middleware(
  request: NextRequest,
  response: NextResponse,
  next: () => void
) {
  console.log("middleware running", request.body);

  response.cookies.set("nextjs", "middleware ran", {
    path: "/",
  });
  return next();
}

export const config = {
  matcher: "/auth/register",
};
