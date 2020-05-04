package main

import (
	"context"
	"fmt"
	"log"

	"github.com/amisini/go-grpc-blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating blog")

	blog := &blogpb.Blog{
		AuthorId: "Admir",
		Title:    "my blog",
		Content:  "My content",
	}

	crtBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("Blog created: %v \n", crtBlogRes)

	blogID := crtBlogRes.GetBlog().GetId()

	fmt.Println("Reading blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "asfrtttt"})
	if err2 != nil {
		fmt.Printf("Error wile reading: %v \n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	rdBlogRes, rdBlogerr := c.ReadBlog(context.Background(), readBlogReq)
	if rdBlogerr != nil {
		fmt.Printf("Error wile reading: %v \n", rdBlogerr)
	}

	fmt.Printf("Blog read: %v \n", rdBlogRes)
}
