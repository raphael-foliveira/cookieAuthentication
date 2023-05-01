import { ChangeEventHandler, FormEventHandler, useState } from "react";

export default function Login() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  const handleChange: ChangeEventHandler<HTMLInputElement> = (e) => {
    setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
    fetch(`http://localhost:8000/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
      credentials: "include",
    })
      .then((response) => {
        console.log(response.status);
        return response.json();
      })
      .then(console.log);
  };

  return (
    <>
      <h1>Login</h1>
      <form action="" onSubmit={handleSubmit}>
        <label htmlFor="username">
          Username
          <input type="text" name="username" onChange={handleChange} />
        </label>
        <label htmlFor="password">
          Password
          <input type="password" name="password" onChange={handleChange} />
        </label>
        <button>Submit</button>
      </form>
    </>
  );
}
