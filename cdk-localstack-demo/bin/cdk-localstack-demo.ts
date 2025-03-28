#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { CdkLocalStackDemoStack } from '../lib/cdk-localstack-demo-stack';

const app = new cdk.App();
new CdkLocalStackDemoStack(app, 'StackOne', { env: { region: 'us-west-2' } });
new CdkLocalStackDemoStack(app, 'StackTwo', { env: { region: 'us-west-2' } });
