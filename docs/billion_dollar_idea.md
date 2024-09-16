# This documentaion contains our billion dollar plan

Hi I am Ankan the founder and probably still the core maintainer of the project. Below is GPT generated project backend structure, probably wrong but as of 16/Sep/2024 this is the future.

Good luck to myself!

# NATS-Based Workflow System Design

## 1. System Overview

In this redesigned system, each block within a workflow operates as a mini NATS client. This approach allows for a more distributed, event-driven architecture where blocks can communicate and coordinate their actions through NATS subjects and messages.

## 2. Core Components

### 2.1 NATS Server
- Acts as the central message broker for the entire system.
- Handles pub-sub and request-reply patterns for inter-block communication.

### 2.2 Workflow Manager
- Parses user-provided JSON configurations.
- Creates and manages Workflow Control Blocks (WCBs).
- Initiates workflow execution by publishing a message to NATS.

### 2.3 Block Executor
- A lightweight runtime that hosts individual blocks.
- Subscribes to NATS subjects for block execution requests.
- Executes block logic and publishes results back to NATS.

### 2.4 Scheduler
- Monitors workflow progress through NATS messages.
- Manages resource allocation and block execution order.
- Handles parallel execution and dependencies between blocks.

## 3. Block Design

Each block is now designed as a mini NATS client with the following characteristics:

```typescript
interface NATSBlock {
  id: string;
  type: 'fetch' | 'html-scraper' | 'hook' | 'response';
  execute: (input: any) => Promise<any>;
  inputSubject: string;
  outputSubject: string;
  errorSubject: string;
  natsClient: NATS.Client;
}
```

## 4. Workflow Execution Flow

1. **Workflow Initialization**:
   - User submits JSON configuration.
   - Workflow Manager creates a WCB and publishes a workflow start message to NATS.

2. **Block Execution**:
   - Each block subscribes to its input subject.
   - When a message is received on the input subject, the block executes its logic.
   - Upon completion, the block publishes its output to the next block's input subject.

3. **Inter-Block Communication**:
   - Blocks communicate exclusively through NATS subjects.
   - The output of one block becomes the input for the next block in the workflow.

4. **Error Handling**:
   - If a block encounters an error, it publishes to its error subject.
   - The Scheduler subscribes to all error subjects and can trigger retries or alternative flows.

5. **Workflow Completion**:
   - The final block in the workflow publishes a completion message.
   - The Workflow Manager receives this message and marks the workflow as complete.

## 5. NATS Subject Structure

- Workflow subjects: `workflow.<workflow_id>.<action>`
  - Example: `workflow.123.start`, `workflow.123.complete`
- Block subjects: `block.<block_id>.<action>`
  - Example: `block.fetch_1.input`, `block.fetch_1.output`, `block.fetch_1.error`

## 6. Advantages of NATS-Based Design

1. **Scalability**: Easy to scale horizontally by adding more Block Executors.
2. **Fault Tolerance**: If a Block Executor fails, another can pick up the work.
3. **Flexibility**: Easy to add new block types or modify existing ones.
4. **Real-time Monitoring**: Use NATS streaming to monitor workflow progress in real-time.
5. **Load Balancing**: NATS can automatically load balance requests across multiple Block Executors.

## 7. Implementation Considerations

1. **NATS Client Libraries**: Use official NATS client libraries for Node.js/TypeScript.
2. **Block Isolation**: Each Block Executor runs in its own process for isolation.
3. **State Management**: Use NATS JetStream for persisting workflow state.
4. **Security**: Implement NATS authentication and encryption for secure communication.
5. **Monitoring**: Utilize NATS monitoring capabilities for system health checks.

## 8. Example: Fetch Block Implementation

```typescript
import * as NATS from 'nats';

class FetchBlock implements NATSBlock {
  id: string;
  type: 'fetch';
  inputSubject: string;
  outputSubject: string;
  errorSubject: string;
  natsClient: NATS.Client;

  constructor(id: string, natsClient: NATS.Client) {
    this.id = id;
    this.type = 'fetch';
    this.inputSubject = `block.${id}.input`;
    this.outputSubject = `block.${id}.output`;
    this.errorSubject = `block.${id}.error`;
    this.natsClient = natsClient;
  }

  async execute(input: any): Promise<any> {
    try {
      const response = await fetch(input.url);
      const data = await response.text();
      return data;
    } catch (error) {
      throw error;
    }
  }

  async start() {
    const subscription = this.natsClient.subscribe(this.inputSubject);
    for await (const msg of subscription) {
      try {
        const input = JSON.parse(msg.data.toString());
        const output = await this.execute(input);
        await this.natsClient.publish(this.outputSubject, JSON.stringify(output));
      } catch (error) {
        await this.natsClient.publish(this.errorSubject, JSON.stringify(error));
      }
    }
  }
}
```

This redesign transforms each block into a mini NATS client, creating a more distributed and event-driven system. Here are some key points about the new design:

- Distributed Architecture: Each block now operates independently as a NATS client, allowing for better scalability and fault tolerance.
- Event-Driven Communication: Blocks communicate through NATS subjects, using publish-subscribe patterns for workflow execution.
- Scalability: The system can easily scale horizontally by adding more Block Executors, as NATS handles load balancing.
- Flexibility: New block types can be added easily by implementing the NATSBlock interface and subscribing to the appropriate subjects.
- Real-time Monitoring: NATS streaming capabilities allow for real-time monitoring of workflow progress.
- Fault Tolerance: If a Block Executor fails, another can pick up the work, improving system reliability.
- State Management: NATS JetStream can be used for persisting workflow state, making the system more robust.

