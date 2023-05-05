import { useEffect, useState } from "react";

interface UserData {
  username: string;
  email: string;
}

export default function GetUser() {
  const [user, setUser] = useState<UserData>();

  useEffect(() => {
    fetch(process.env.NEXT_PUBLIC_API_URL + "/user", {
      method: "GET",
      credentials: "include",
    })
      .then((response) => response.json())
      .then((data) => setUser(data));
  }, []);

  return (
    <>
      <h1>User</h1>
      <h3>Username: {user && user.username}</h3>
      <h3>Email: {user && user.email}</h3>
    </>
  );
}
