import { useEffect, useState } from "react";

export default function GetUser() {
  const [user, setUser] = useState({
    username: "",
    email: "",
  });

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`http://localhost:8000/user`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Cookie: document.cookie,
        },
        credentials: "include",
      });
      const data = await response.json();
      setUser(data);
    };
    fetchData();
  }, []);

  return (
    <>
      <h1>User</h1>
      <h3>Username: {user.username}</h3>
      <h3>Email: {user.email}</h3>
    </>
  );
}
