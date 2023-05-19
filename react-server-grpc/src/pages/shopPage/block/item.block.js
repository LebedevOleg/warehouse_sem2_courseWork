import {
	Box,
	Button,
	Grid,
	Stack,
	TableCell,
	TableRow,
	Typography,
} from "@mui/material";

import React, { useContext } from "react";

const ItemBlock = (props) => {
	const handleClick = () => {
		const backet = JSON.parse(localStorage.getItem("backet"));
		backet.push(props.product);
		localStorage.setItem("backet", JSON.stringify(backet));
	};

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
				{props.product.c_name}; Единица продажи: {props.product.dim}
			</TableCell>
			<TableCell align="center">
				Цена за единицу: {props.product.pfu}
			</TableCell>
			<TableCell align="center">
				<Button onClick={handleClick}>Добавить в корзину</Button>
			</TableCell>
		</TableRow>
	);
};

export default ItemBlock;

{
	/* <Grid>
				<Typography>{props.name}</Typography>
				<br />
				<Typography>{props.id}</Typography>
			</Grid>
			<Grid>
				<Typography>
					Категория: {props.c_name}; Единица продажи: {props.dim}{" "}
				</Typography>
			</Grid>
			<Grid>
				<Typography>Цена за единицу: {props.pfu}</Typography>
			</Grid> */
}
