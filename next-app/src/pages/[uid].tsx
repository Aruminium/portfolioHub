import { useUser } from "lib/auth";
import type { NextPage } from "next";
import Router, { useRouter } from "next/router";
import { useEffect } from "react";

const User: NextPage = () => {
  const router = useRouter();
  const user = useUser();
  const { uid } = router.query;

  useEffect(() => {
    async () => {
      if (router.isReady) {
        //dbからユーザ情報を持ってくる
        try {
        } catch (err) {
          console.error(err);
          //失敗したら404返す
        }
      }
    };
  });

  return (
    <div>
      <h1>user</h1>
    </div>
  );
};

export default User;
