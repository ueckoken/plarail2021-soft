import networkx as nx
import matplotlib.pyplot as plt
import yaml

with open("./hallsensor.yaml") as f:
    y = yaml.safe_load(f)

y = y["halls"]

G = nx.DiGraph()

G.add_nodes_from(list(map(lambda x: x["name"], y)))
s = []
for i, n in enumerate(y):
    print("name")
    if n["nexts"]==None:
        continue
    for j, m in enumerate(n["nexts"]):
        print(m)
        s.append((n["name"],m))
G.add_edges_from(s)
plt.figure(figsize=(15,15))
nx.spring_layout(G)
nx.draw(G,node_color="w",alpha=0.6, node_size=3, with_labels=True)
plt.show()
