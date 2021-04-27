package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bndr/gojenkins"
)

func main() {
	ctx := context.Background()
	jenkins, err := gojenkins.CreateJenkins(nil, "http://192.168.56.102:8080", "admin", "88888888").Init(ctx)

	if err != nil {
		log.Printf("Jenkins connect error - %s\n", err)
		return
	} else {
		log.Printf("Jenkins connect Success \n")
	}

	// 获取节点信息
	nodes, _ := jenkins.GetAllNodes(ctx)

	// 检查节点的状态
	for _, node := range nodes {
		fmt.Println(node.IsOnline(ctx))
		fmt.Println(node)
	}

	// 创建文件夹 参数1： go上下文， 参数2： 需要创建的文件夹  参数3： 父文件夹
	folder, err := jenkins.CreateFolder(ctx, "go-server", "lixiong")
	if err != nil {
		fmt.Println("创建文件夹失败", err)
		// return
	} else {
		fmt.Println("文件夹创建成功", folder)
	}
	ok, err := jenkins.DeleteJob(ctx, "testgo")
	if err != nil {
		fmt.Println("删除异常", err)
	} else {
		fmt.Println("删除成功", ok)
	}

	configString := `<?xml version='1.1' encoding='UTF-8'?>
	<flow-definition plugin="workflow-job@2.40">
	  <actions>
		<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.7.2"/>
		<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.7.2">
		  <jobProperties/>
		  <triggers/>
		  <parameters/>
		  <options/>
		</org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
	  </actions>
	  <description></description>
	  <keepDependencies>false</keepDependencies>
	  <properties>
		<hudson.model.ParametersDefinitionProperty>
		  <parameterDefinitions>
			<hudson.model.StringParameterDefinition>
			  <name>deployServices</name>
			  <description>deployServices</description>
			  <defaultValue>web</defaultValue>
			  <trim>true</trim>
			</hudson.model.StringParameterDefinition>
			<hudson.model.ChoiceParameterDefinition>
			  <name>BuildType</name>
			  <description>build type case</description>
			  <choices class="java.util.Arrays$ArrayList">
				<a class="string-array">
				  <string>mvn</string>
				  <string>gradle</string>
				  <string>npm</string>
				  <string>ant</string>
				</a>
			  </choices>
			</hudson.model.ChoiceParameterDefinition>
			<hudson.model.ChoiceParameterDefinition>
			  <name>BuildShell</name>
			  <description>build shell</description>
			  <choices class="java.util.Arrays$ArrayList">
				<a class="string-array">
				  <string>-v</string>
				  <string>npm install</string>
				  <string>npm run build</string>
				  <string>build --info --debug</string>
				  <string>install &amp;&amp; npm build</string>
				</a>
			  </choices>
			</hudson.model.ChoiceParameterDefinition>
			<hudson.model.StringParameterDefinition>
			  <name>dest</name>
			  <description>/test/</description>
			  <defaultValue>root@192.168.56.101:/test/</defaultValue>
			  <trim>true</trim>
			</hudson.model.StringParameterDefinition>
			<hudson.model.StringParameterDefinition>
			  <name>src</name>
			  <description>gradle_study-1.0-SNAPSHOT.jar</description>
			  <defaultValue>/root/.jenkins/workspace/pipeline-gradle/build/libs/gradle_study-1.0-SNAPSHOT.jar</defaultValue>
			  <trim>true</trim>
			</hudson.model.StringParameterDefinition>
			<hudson.model.StringParameterDefinition>
			  <name>module</name>
			  <description>shell</description>
			  <defaultValue>shell</defaultValue>
			  <trim>true</trim>
			</hudson.model.StringParameterDefinition>
			<hudson.model.StringParameterDefinition>
			  <name>args</name>
			  <description>bash /test/start.sh</description>
			  <defaultValue>bash /test/start.sh</defaultValue>
			  <trim>true</trim>
			</hudson.model.StringParameterDefinition>
		  </parameterDefinitions>
		</hudson.model.ParametersDefinitionProperty>
	  </properties>
	  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.86">
		<scm class="hudson.plugins.git.GitSCM" plugin="git@4.4.5">
		  <configVersion>2</configVersion>
		  <userRemoteConfigs>
			<hudson.plugins.git.UserRemoteConfig>
			  <url>http://192.168.56.101:8888/root/jenkinslib.git</url>
			  <credentialsId>f9c9f149-0c46-4cf2-9586-37f39af51bcb</credentialsId>
			</hudson.plugins.git.UserRemoteConfig>
		  </userRemoteConfigs>
		  <branches>
			<hudson.plugins.git.BranchSpec>
			  <name>*/master</name>
			</hudson.plugins.git.BranchSpec>
		  </branches>
		  <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
		  <submoduleCfg class="list"/>
		  <extensions/>
		</scm>
		<scriptPath>Jenkinsfile</scriptPath>
		<lightweight>true</lightweight>
	  </definition>
	  <triggers/>
	  <disabled>false</disabled>
	</flow-definition>
	`
	// 创建任务  参数1： jenkins上下文， 参数2： 创建参数模板， 参数3： 创建的job名称
	job, err := jenkins.CreateJob(ctx, configString, "testgo")
	if err != nil {
		fmt.Println("job创建失败", err)
		return
	}
	fmt.Println("job创建成功", job)

	// 在指定文件夹下创建job
	fjob, err := jenkins.CreateJobInFolder(ctx, configString, "lx", "gojenkins", "lixiong")
	if err != nil {
		fmt.Println("创建失败", err)
	} else {
		fmt.Println("创建成功", fjob)
	}

	// 构建任务 任务执行
	ARGS := map[string]string{}
	resultcode, err := jenkins.BuildJob(ctx, "testgo", map[string]string{})
	if err != nil {
		fmt.Println("执行失败：", err)
	} else {
		fmt.Println(resultcode, ARGS)
	}

	// 获取任务
	jobname, err := jenkins.GetJob(ctx, "testgo")
	if err != nil {
		fmt.Println("获取任务失败", err)
	} else {
		fmt.Println("获取任务成功:", jobname)
	}

	buildName, err := jenkins.GetBuild(ctx, "testgo", 22)
	if err != nil {
		fmt.Println("获取指定任务失败", err)
	} else {
		fmt.Println("获取指定成功:", buildName)
	}
}
