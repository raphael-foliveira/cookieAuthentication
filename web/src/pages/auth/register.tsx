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
    fetch(process.env.NEXT_PUBLIC_API_URL + "/register", {
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
          <div className={styles.formElement}>
            <label htmlFor="email">Email</label>
            <input type="text" name="email" onChange={handleChange} />
          </div>
          <div className={styles.formElement}>
            <label htmlFor="username">Username</label>
            <input type="text" name="username" onChange={handleChange} />
          </div>
          <div className={styles.formElement}>
            <label htmlFor="password">Password</label>
            <input type="text" name="password" onChange={handleChange} />
          </div>

          <div className={styles.formElement}>
            <button>Submit</button>
          </div>
        </div>
      </form>
    </>
  );
}
