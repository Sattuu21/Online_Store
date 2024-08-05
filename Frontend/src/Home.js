import * as React from "react";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Button from "@mui/material/Button";
import { Link as RouterLink } from "react-router-dom";

const products = [
    {
        title: "Forever Pants",
        image: process.env.PUBLIC_URL + "/images/pants.jpg",
        description: "description",
        link: "/product",     },
      {
        title: "Forever Shirt",
        image: process.env.PUBLIC_URL + "/images/shirt.jpg",
        description: "description",
        link: "/product",},
      {
        title: "Forever Shorts",
        image: process.env.PUBLIC_URL + "/images/shorts.jpg",
        description: "description",
        link: "/product",},
];

export default function Home() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <Typography variant="h2" align="center" gutterBottom>
        Welcome to Our Online Store
      </Typography>
      <Grid container spacing={4}>
        {products.map((product, index) => (
          <Grid item xs={12} sm={6} md={4} key={index}>
            <Card>
              <CardMedia
                component="img"
                height="140"
                image={product.image}
                alt={product.title}
              />
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {product.title}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {product.description}
                </Typography>
              </CardContent>
              <Button size="small" component={RouterLink} to={product.link}>
                View Product
              </Button>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
}
