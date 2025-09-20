import { Outlet } from "react-router-dom";
import Header from "./components/pages/Header.tsx";

const Layout = () => {
  return (
    <div className="flex-1 min-w-screen min-h-screen">
      <Header />
      <main>
        <Outlet />
      </main>
    </div>
  );
};

export default Layout;
