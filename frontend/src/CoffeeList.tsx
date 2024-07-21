import React from "react"
import { Table } from "react-bootstrap";
import axios from "axios";

type Product = {
    name: string,
    price: string,
    sku: string,
}

const CoffeeList = () => {
    const [products, setProducts] = React.useState<Product[]>([])
    React.useEffect(() => {
        const fetchProducts = async () => {
            //axios.get(window.globalThis.api)
        }
    }, [])
    return (
        <div>
            <h1 style={{ marginBottom: "40px" }}>Menu</h1>
            <Table>
                <thead>
                    <tr>
                        <th>
                            Name
                        </th>
                        <th>
                            Price
                        </th>
                        <th>
                            SKU
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {<RenderProducts products={products} />}
                </tbody>
            </Table>
        </div>
    )
}

interface RenderProductsParams {
    products: Product[]
}

const RenderProducts: React.FC<RenderProductsParams> = ({ products }) => {
    return (
        products.map((p, index) => {
            return (
                <tr key={index}>
                    <td>{p.name}</td>
                    <td>{p.price}</td>
                    <td>{p.sku}</td>
                </tr>
            )

        })
    )
}

export default CoffeeList;