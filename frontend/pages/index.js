import { useEffect, useState } from "react";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";

export default function Home() {
  const [status, setStatus] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const controller = new AbortController();

    fetch(`${API_BASE_URL}/healthz`, { signal: controller.signal })
      .then(async (response) => {
        if (!response.ok) {
          throw new Error(`API error: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => setStatus(data.status))
      .catch((err) => setError(err.message));

    return () => controller.abort();
  }, []);

  return (
    <main>
      <h1>Go x Next.js サンプル</h1>
      <p>バックエンドのヘルスチェック `/healthz` の結果を表示します。</p>
      {status && <p>status: <strong>{status}</strong></p>}
      {error && <p style={{ color: "#f87171" }}>error: {error}</p>}
      {!status && !error && <p>読み込み中...</p>}
    </main>
  );
}
