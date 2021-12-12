package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		args := &lambda.FunctionArgs{
			Handler: pulumi.String("main"),
			Runtime: pulumi.String("go1.x"),
			Code:    pulumi.NewFileArchive("./handler/main.zip"),
		}

		function, err := lambda.NewFunction(
			ctx, "MyLambda", args, pulumi.DependsOn([]pulumi.Resource{}),
		)
		if err != nil {
			return err
		}

		ctx.Export("lambda", function.Arn)
		return nil
	})
}
