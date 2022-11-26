import { login, logout, useUser } from "lib/auth";
import type { NextPage } from "next";
import Head from "next/head";
import Dashboard from "components/dashboard/Dashboard";

const Home: NextPage = () => {
  const user = useUser();

  const handleLogin = (): void => {
    login().catch((error) => console.error(error));
  };

  const handleLogout = (): void => {
    logout().catch((error) => console.error(error));
  };

  return (
    <div>
      {user !== null ? (
        <div>
          <Dashboard />
        </div>
      ) : (
        <div>
          <h2>ログインしていない</h2>
          <button onClick={handleLogin}>ログイン</button>
        </div>
      )}
    </div>
  );
};

export default Home;
