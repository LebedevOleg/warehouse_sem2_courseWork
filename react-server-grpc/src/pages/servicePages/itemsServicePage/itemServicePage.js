import axios from "axios";
import React, { useCallback, useEffect, useState } from "react";
import {
	TableContainer,
	Paper,
	TableHead,
	TableRow,
	TableCell,
	Table,
	TableBody,
	Typography,
} from "@mui/material";
import ItemServiceBlock from "./blocks/item.block";

const ItemServicePage = () => {
	const [products, setProducts] = useState([]);

	const GetAllProducts = useCallback(async () => {
		await axios.get("http://localhost:8000/getallitems").then((res) => {
			setProducts(res.data.allItems);
		});
	}, []);

	useEffect(() => {
		GetAllProducts();
	}, [GetAllProducts]);
	return (
		<>
			<Typography variant="h3">Редактирование товаров </Typography>
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
							<ItemServiceBlock product={product} />
						))}
					</TableBody>
				</Table>
			</TableContainer>
		</>
	);
};

export default ItemServicePage;
