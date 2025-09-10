import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";
import dataProvider from "./dataProvider.ts";
import ProductList from "./components/product/ProductList";

const App = () => (
  <Admin layout={Layout} dataProvider={dataProvider} disableTelemetry>
    <Resource name="products" list={ProductList} />
  </Admin>
);

export default App;
