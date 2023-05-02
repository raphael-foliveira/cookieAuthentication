import { useRouter } from "next/router";
import { ChangeEventHandler, FormEventHandler, useState } from "react";
import styles from "./auth.module.css";

interface RegisterFormData {
  username: string;
  email: string;
  password: string;
}

export default function Register() {
  const [formData, setFormData] = useState<RegisterFormData>({
    username: "",
    email: "",
    password: "",
  });

  const router = useRouter();

  const handleChange: ChangeEventHandler<HTMLInputElement> = (e) => {
    setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
    fetch("http://localhost:8000/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
      credentials: "include",
    }).then((response) => {
      console.log(response.status);
      if (response.status === 201) {
        router.push("/auth/login");
      }
      return response.json();
    });
  };

  return (
    <>
      <h1>Register</h1>
      <form action="" onSubmit={handleSubmit}>
        <div className={styles.formContainer}>
          <label htmlFor="email">
            <p>Email</p>
            <input type="text" name="email" onChange={handleChange} />
          </label>
          <label htmlFor="username">
            <p>Username</p>
            <input type="text" name="username" onChange={handleChange} />
          </label>
          <label htmlFor="password">
            <p>Password</p>
            <input type="text" name="password" onChange={handleChange} />
          </label>
          <p>
            <button>Submit</button>
          </p>
        </div>
      </form>
    </>
  );
}
