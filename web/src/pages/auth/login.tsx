import { useRouter } from "next/router";
import { ChangeEventHandler, FormEventHandler, useState } from "react";
import styles from "./auth.module.css";

interface LoginFormData {
  username: string;
  password: string;
}

export default function Login() {
  const [formData, setFormData] = useState<LoginFormData>({
    username: "",
    password: "",
  });
  const router = useRouter();

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
    }).then((response) => {
      console.log(response.status);
      if (response.status === 200) {
        router.push("/auth/get-user");
      }
      return response.json();
    });
  };

  return (
    <>
      <h1>Login</h1>
      <form action="" onSubmit={handleSubmit}>
        <div className={styles.formContainer}>
          <p>
            <label htmlFor="username">Username</label>
            <input type="text" name="username" onChange={handleChange} />
          </p>
          <p>
            <label htmlFor="password">Password</label>
            <input type="password" name="password" onChange={handleChange} />
          </p>
          <p>
            <button>Submit</button>
          </p>
        </div>
      </form>
    </>
  );
}
