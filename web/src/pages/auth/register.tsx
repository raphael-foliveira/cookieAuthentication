import { ChangeEventHandler, FormEventHandler, useState } from "react";

interface RegisterFormData {
  username: string;
  email: string;
  password: string;
}

export default function Register() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
  });

  const handleChange: ChangeEventHandler<HTMLInputElement> = (e) => {
    setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
    fetch("http://localhost:8000/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    }).then((response) => {
      console.log(response.status);
      return response.json();
    });
  };

  return (
    <>
      <h1>Register</h1>
      <form action="" onSubmit={handleSubmit}>
        <label htmlFor="email">
          <input type="text" name="email" onChange={handleChange} />
        </label>
        <label htmlFor="username">
          <input type="text" name="username" onChange={handleChange} />
        </label>
        <label htmlFor="password">
          <input type="text" name="password" onChange={handleChange} />
        </label>
        <button>Submit</button>
      </form>
    </>
  );
}
