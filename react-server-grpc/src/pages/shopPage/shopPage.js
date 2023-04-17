import {
	TableContainer,
	Paper,
	TableHead,
	TableRow,
	TableCell,
	Table,
	TableBody,
} from "@mui/material";
import axios from "axios";
import React, { useCallback } from "react";
import ItemBlock from "./block/item.block";

const ShopPage = () => {
	const [products, setProducts] = React.useState([]);

	const GetAllProducts = useCallback(async () => {
		await axios.get("http://localhost:8000/getall").then((res) => {
			setProducts(res.data.allItems);
		});
	}, []);

	React.useEffect(() => {
		GetAllProducts();
	}, [GetAllProducts]);

	return (
		<div>
			<h1>Shop Page</h1>
			<TableContainer component={Paper}>
				<Table
					sx={{ minWidth: 500 }}
					padding="none"
					aria-label="simple table"
				>
					<TableHead>
						<TableRow>
							<TableCell>Product Name</TableCell>
							<TableCell align="center">Tags</TableCell>
							<TableCell align="center">Price</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{products.map((product) => (
							<ItemBlock product={product} />
						))}
					</TableBody>
				</Table>
			</TableContainer>
		</div>
	);
};

export default ShopPage;
