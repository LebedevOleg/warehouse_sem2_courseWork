import {
	Box,
	Button,
	Grid,
	Stack,
	TableCell,
	TableRow,
	Typography,
} from "@mui/material";

import React from "react";
import StockChangeModal from "../modal/stockChange.modal";

const StockServiceBlock = (props) => {
	return (
		<TableRow key={props.product.id}>
			<TableCell align="center">
				<Stack direction="row" spacing={3}>
					<Typography variant="subtitle2">
						{props.product.id}
					</Typography>
					<Typography variant="h6">{props.product.name}</Typography>
				</Stack>
			</TableCell>
			<TableCell align="center">
				Всего товаров на складе:{props.product.items_count}
			</TableCell>
			<TableCell align="center">Адрес {props.product.address}</TableCell>
			<TableCell align="center">
				<StockChangeModal stock={props.product} />
				<Button>Удалить</Button>
			</TableCell>
		</TableRow>
	);
};

export default StockServiceBlock;
