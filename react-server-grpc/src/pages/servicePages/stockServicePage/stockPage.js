import {
	Box,
	Paper,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	Typography,
} from "@mui/material";
import React, { useCallback, useEffect } from "react";
import { useAuth } from "../../../hooks/auth.hook";
import StockServiceBlock from "./blocks/stock.block";

const StockPage = () => {
	const { token } = useAuth();
	const [stocks, setStocks] = React.useState([]);

	const GetAllStocks = useCallback(async () => {}, []);

	useEffect(() => {
		GetAllStocks();
	}, [GetAllStocks]);

	return (
		<>
			<Typography variant="h3">Управление Складами</Typography>
			<Box>
				<TableContainer component={Paper}>
					<Table
						sx={{ minWidth: 500 }}
						padding="none"
						aria-label="simple table"
					>
						<TableHead>
							<TableRow>
								<TableCell>Stock Name</TableCell>
								<TableCell align="center">
									Items count
								</TableCell>
								<TableCell align="center">Price</TableCell>
							</TableRow>
						</TableHead>
						<TableBody>
							{stocks.map((product) => (
								<StockServiceBlock product={product} />
							))}
						</TableBody>
					</Table>
				</TableContainer>
			</Box>
		</>
	);
};
