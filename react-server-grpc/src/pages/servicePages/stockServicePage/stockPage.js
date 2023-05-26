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
import axios from "axios";
import React, { useCallback, useContext, useEffect } from "react";
import { AuthContext } from "../../../context/auth.context";
import StockServiceBlock from "./blocks/stock.block";

const StockPage = () => {
	const auth = useContext(AuthContext);
	const [stocks, setStocks] = React.useState([]);

	const GetAllStocks = useCallback(async () => {
		await axios
			.get("http://localhost:8000/getallstocks", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setStocks(res.data.stocks);
			});
	}, []);

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

export default StockPage;
